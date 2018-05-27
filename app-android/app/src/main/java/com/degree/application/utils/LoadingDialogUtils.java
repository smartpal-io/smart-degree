package com.degree.application.utils;

import android.app.Activity;
import android.app.ProgressDialog;

public class LoadingDialogUtils {

    private static ProgressDialog progressDialog;

    public static void displayLoadingDialog(Activity activity) {
        progressDialog = new ProgressDialog(activity);
        progressDialog.setTitle("Loading");
        progressDialog.setMessage("Wait while loading...");
        progressDialog.setCancelable(false); // disable dismiss by tapping outside of the dialog
        progressDialog.show();
    }

    public static void hideLoadingDialog(){
        if(progressDialog!=null) {
            progressDialog.hide();
            progressDialog = null;
        }
    }
}
