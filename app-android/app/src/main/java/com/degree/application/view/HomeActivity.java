package com.degree.application.view;

import android.animation.Animator;
import android.animation.AnimatorListenerAdapter;
import android.annotation.TargetApi;
import android.support.annotation.UiThread;
import android.support.v7.app.AppCompatActivity;
import android.app.LoaderManager.LoaderCallbacks;

import android.content.CursorLoader;
import android.content.Loader;
import android.database.Cursor;
import android.net.Uri;
import android.os.AsyncTask;

import android.os.Build;
import android.os.Bundle;
import android.provider.ContactsContract;
import android.util.Log;
import android.view.KeyEvent;
import android.view.View;
import android.view.View.OnClickListener;
import android.view.inputmethod.EditorInfo;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import com.degree.application.BuildConfig;
import com.degree.application.R;
import com.degree.application.contract.SmartDegree;

import org.apache.commons.codec.binary.Hex;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.abi.datatypes.generated.Bytes32;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jFactory;
import org.web3j.protocol.http.HttpService;
import org.web3j.tx.Contract;
import org.web3j.tx.ManagedTransaction;
import org.web3j.tx.ReadonlyTransactionManager;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

import rx.plugins.RxJavaHooks;

import static org.apache.commons.codec.binary.Hex.decodeHex;

public class HomeActivity extends AppCompatActivity {


    // UI references.
    private EditText verifyDegreeId;
    private EditText verifyHash;
    private View mProgressView;
    private View degreeFormView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_home);

        verifyDegreeId = findViewById(R.id.degreeId);

        verifyHash = findViewById(R.id.hash);
        verifyHash.setOnEditorActionListener((textView, id, keyEvent)
                -> id == EditorInfo.IME_ACTION_DONE || id == EditorInfo.IME_NULL);

        Button verifyButton = findViewById(R.id.verify);
        verifyButton.setOnClickListener(view -> verify());

        degreeFormView = findViewById(R.id.degree_form);
        mProgressView = findViewById(R.id.verify_progress);
    }


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

                        Boolean result = contract.verify(
                                new BigInteger(verifyDegreeId.getText().toString()),
                                decodeHex(verifyHash.getText().toString().toCharArray())).send();

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


}

