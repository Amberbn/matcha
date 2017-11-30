package io.gomatcha.matcha;

import android.app.Activity;
import android.content.Context;
import android.os.Handler;
import android.text.Editable;
import android.text.InputType;
import android.text.SpannableString;
import android.text.SpannableStringBuilder;
import android.text.TextWatcher;
import android.util.Log;
import android.view.Gravity;
import android.view.KeyEvent;
import android.view.View;
import android.view.inputmethod.EditorInfo;
import android.view.inputmethod.InputMethodManager;
import android.widget.EditText;
import android.widget.TextView;

import com.google.protobuf.InvalidProtocolBufferException;

import io.gomatcha.bridge.GoValue;
import io.gomatcha.matcha.proto.text.PbText;
import io.gomatcha.matcha.proto.view.PbTextInput;

class MatchaTextInputView extends MatchaChildView {
    EditText view;
    boolean editing;
    boolean focused;
    MatchaViewNode viewNode;

    static {
        MatchaView.registerView("gomatcha.io/matcha/view/textinput", new MatchaView.ViewFactory() {
            @Override
            public MatchaChildView createView(Context context, MatchaViewNode node) {
                return new MatchaTextInputView(context, node);
            }
        });
    }

    public MatchaTextInputView(Context context, MatchaViewNode node) {
        super(context);
        viewNode = node;

        view = new EditText(context);
        view.setPadding(0, 0, 0, 0);
        view.setBackground(null);
        view.setGravity(Gravity.TOP);
        view.setOnEditorActionListener(new TextView.OnEditorActionListener() {
            @Override
            public boolean onEditorAction(TextView textView, int i, KeyEvent keyEvent) {
                boolean handled = false;
                if (i == EditorInfo.IME_ACTION_DONE) {
                    Log.v("x", "submit");
                    MatchaTextInputView.this.viewNode.call("OnSubmit");
                    handled = true;
                }
                return handled;
            }
        });
        view.addTextChangedListener(new TextWatcher() {
            @Override
            public void beforeTextChanged(CharSequence charSequence, int i, int i1, int i2) {
                // no-op
            }
            @Override
            public void onTextChanged(CharSequence charSequence, int i, int i1, int i2) {
                if (!editing) {
                    PbText.StyledText styledText = Protobuf.toProtobuf((SpannableStringBuilder) charSequence);
                    PbTextInput.TextInputEvent proto = PbTextInput.TextInputEvent.newBuilder().setStyledText(styledText).build();
                    MatchaTextInputView.this.viewNode.call("OnTextChange", new GoValue(proto.toByteArray()));
                }
            }
            @Override
            public void afterTextChanged(Editable editable) {
                // no-op
            }
        });
        view.setOnFocusChangeListener(new OnFocusChangeListener() {
            @Override
            public void onFocusChange(View view, boolean b) {
                if (!editing) {
                    PbTextInput.TextInputFocusEvent proto = PbTextInput.TextInputFocusEvent.newBuilder().setFocused(b).build();
                    MatchaTextInputView.this.viewNode.call("OnFocus", new GoValue(proto.toByteArray()));
                }
            }
        });
        addView(view);
    }

    @Override
    public void setNativeState(byte[] nativeState) {
        super.setNativeState(nativeState);
        try {
            PbTextInput.TextInput proto = PbTextInput.TextInput.parseFrom(nativeState);
            editing = true;
            SpannableString str = Protobuf.newAttributedString(proto.getStyledText());
            if (!str.toString().equals(view.getText().toString())) {
                view.setText(str, TextView.BufferType.SPANNABLE);
            }
            if (view.hasFocus() && !proto.getFocused()) {
                view.clearFocus();
                InputMethodManager imm = (InputMethodManager) MatchaTextInputView.this.getContext().getSystemService(Context.INPUT_METHOD_SERVICE);
                imm.hideSoftInputFromWindow(((Activity)this.getContext()).getCurrentFocus().getWindowToken(), InputMethodManager.HIDE_NOT_ALWAYS);
            } else if (!view.hasFocus() && proto.getFocused()) {
                new Handler().postDelayed(new Runnable() {
                    public void run() {
                        editing = true;
                        view.requestFocus();
                        InputMethodManager imm = (InputMethodManager) MatchaTextInputView.this.getContext().getSystemService(Context.INPUT_METHOD_SERVICE);
                        imm.showSoftInput(view, 0);
                        editing = false;
                    }
                }, 100);
            }
            int inputType;
            switch (proto.getKeyboardType()) {
                case TEXT_TYPE:
                    inputType = InputType.TYPE_CLASS_TEXT;
                break;
                case NUMBER_TYPE:
                    inputType = InputType.TYPE_CLASS_NUMBER;
                break;
                case PHONE_TYPE:
                    inputType = InputType.TYPE_CLASS_PHONE;
                break;
                case EMAIL_TYPE:
                    inputType = InputType.TYPE_CLASS_TEXT | InputType.TYPE_TEXT_VARIATION_EMAIL_ADDRESS;
                break;
                case URL_TYPE:
                    inputType = InputType.TYPE_CLASS_TEXT | InputType.TYPE_TEXT_VARIATION_URI;
                break;
                case DATE_TIME_TYPE:
                    inputType = InputType.TYPE_CLASS_DATETIME | InputType.TYPE_DATETIME_VARIATION_NORMAL;
                break;
                default:
                    inputType = InputType.TYPE_CLASS_TEXT;
            }
            if (proto.getSecureTextEntry()) {
                inputType |= InputType.TYPE_TEXT_VARIATION_PASSWORD;
            }

            int imeOptions = 0;
            switch (proto.getKeyboardReturnType()) {
                case DEFAULT_RETURN_TYPE:
                break;
                case GO_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_GO;
                break;
                case GOOGLE_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_GO;
                break;
                case JOIN_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_DONE;
                break;
                case NEXT_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_NEXT;
                break;
                case ROUTE_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_DONE;
                break;
                case SEARCH_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_SEARCH;
                break;
                case SEND_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_SEND;
                break;
                case YAHOO_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_DONE;
                break;
                case DONE_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_DONE;
                break;
                case EMERGENCY_CALL_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_DONE;
                break;
                case CONTINUE_RETURN_TYPE:
                    imeOptions = EditorInfo.IME_ACTION_DONE;
                break;
                default:
            }
            view.setImeOptions(imeOptions);
            view.setSingleLine(proto.getMaxLines() == 1);

            view.setHint(Protobuf.newAttributedString(proto.getPlaceholderText()));
            focused = proto.getFocused();
            editing = false;
            
        } catch (InvalidProtocolBufferException e) {
        }
    }
}
