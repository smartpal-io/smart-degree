package com.degree.application.view;

import android.animation.Animator;
import android.animation.AnimatorListenerAdapter;
import android.content.Intent;
import android.os.Bundle;
import android.support.annotation.UiThread;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import com.degree.application.BuildConfig;
import com.degree.application.R;
import com.degree.application.contract.SmartDegree;
import com.google.zxing.client.android.CaptureActivity;

import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jFactory;
import org.web3j.protocol.http.HttpService;
import org.web3j.tx.Contract;
import org.web3j.tx.ManagedTransaction;
import org.web3j.tx.ReadonlyTransactionManager;

import java.math.BigInteger;

import static org.web3j.crypto.Hash.sha3;
import static org.web3j.utils.Numeric.hexStringToByteArray;
import static org.web3j.utils.Numeric.toHexString;

public class HomeActivity extends AppCompatActivity {


    // UI references.
    private EditText verifyDegreeId;
    private EditText studentName, degreeLabel;
    private View mProgressView;
    private View degreeFormView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_home);

        verifyDegreeId = findViewById(R.id.degreeId);

        studentName = findViewById(R.id.studentName);
        degreeLabel = findViewById(R.id.degreeLabel);

        findViewById(R.id.verify).setOnClickListener(view -> verify());

        findViewById(R.id.scan).setOnClickListener(view -> scan());

        degreeFormView = findViewById(R.id.degree_form);
        mProgressView = findViewById(R.id.verify_progress);
    }


    /**
     * Call SmartDegree contract in order to verify the validity of a degree
     */
    private void verify() {
        new Thread(() -> {
            try {
                showProgress(true);
                Web3j web3 = Web3jFactory.build(new HttpService(BuildConfig.ETH_NETWORK_URL));
                web3.web3ClientVersion().observable().subscribe(x -> {
                    try {
                        SmartDegree contract = SmartDegree.load(
                                SmartDegree.getPreviouslyDeployedAddress(BuildConfig.ETH_NETWORK_ID),
                                web3,
                                new ReadonlyTransactionManager(web3, ""),
                                ManagedTransaction.GAS_PRICE, Contract.GAS_LIMIT);

                        String degreeId = verifyDegreeId.getText().toString();
                        String studentNameValue = studentName.getText().toString();
                        String degreeLabelValue = degreeLabel.getText().toString();
                        byte[] verifyHash = sha3(degreeId.concat(studentNameValue.concat(degreeLabelValue)).getBytes());

                        Boolean result = contract.verify(
                                new BigInteger(degreeId),
                                verifyHash
                        ).send();

                        displayMessage("-> " + result);

                    } catch (Exception e) {
                        displayMessage(e.getMessage());
                    }
                });
            } catch (Exception e) {
                displayMessage(e.getMessage());
            }
        }).start();
    }

    /**
     * Start activity in order to scan a QR Code
     */
    public void scan() {
        Intent intent = new Intent(getApplicationContext(), CaptureActivity.class);
        intent.setAction("com.google.zxing.client.android.SCAN");
        intent.putExtra("SAVE_HISTORY", false);
        startActivityForResult(intent, 0);
    }

    /**
     * Shows a message and stop the progress UI.
     */
    @UiThread
    private void displayMessage(String message) {
        runOnUiThread(() -> {
            Toast.makeText(getApplicationContext(), "Verify result " + message, Toast.LENGTH_LONG).show();
            showProgress(false);
        });
    }

    /**
     * Shows the progress UI and hides the verify form.
     */
    @UiThread
    private void showProgress(final boolean show) {
        runOnUiThread(() -> {
            int shortAnimTime = getResources().getInteger(android.R.integer.config_shortAnimTime);

            degreeFormView.setVisibility(show ? View.GONE : View.VISIBLE);
            degreeFormView.animate().setDuration(shortAnimTime).alpha(
                    show ? 0 : 1).setListener(new AnimatorListenerAdapter() {
                @Override
                public void onAnimationEnd(Animator animation) {
                    degreeFormView.setVisibility(show ? View.GONE : View.VISIBLE);
                }
            });

            mProgressView.setVisibility(show ? View.VISIBLE : View.GONE);
            mProgressView.animate().setDuration(shortAnimTime).alpha(
                    show ? 1 : 0).setListener(new AnimatorListenerAdapter() {
                @Override
                public void onAnimationEnd(Animator animation) {
                    mProgressView.setVisibility(show ? View.VISIBLE : View.GONE);
                }
            });
        });
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        if (requestCode == 0) {
            if (resultCode == RESULT_OK) {
                String contents = data.getStringExtra("SCAN_RESULT");
                if(contents!=null && !contents.isEmpty()) {
                    // format of qr code message
                    // degree_id/student_name/degree_label
                    String[] strings = contents.split("/");
                    if (strings.length == 3) {
                        verifyDegreeId.setText(strings[0]);
                        studentName.setText(strings[1]);
                        degreeLabel.setText(strings[2]);
                    }
                }
            }
        }
    }

}

