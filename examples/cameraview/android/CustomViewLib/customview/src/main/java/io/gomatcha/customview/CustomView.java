package io.gomatcha.customview;

import android.content.Context;
import android.support.v7.widget.SwitchCompat;
import android.util.DisplayMetrics;
import android.widget.CompoundButton;

import com.google.protobuf.InvalidProtocolBufferException;

import io.gomatcha.bridge.GoValue;
import io.gomatcha.customview.proto.CustomViewProto;
import io.gomatcha.matcha.MatchaChildView;
import io.gomatcha.matcha.MatchaView;
import io.gomatcha.matcha.MatchaViewNode;

public class CustomView extends MatchaChildView {
    MatchaViewNode viewNode;
    SwitchCompat view;
    boolean checked;

    static {
        MatchaView.registerView("github.com/overcyn/customview", new MatchaView.ViewFactory() {
            @Override
            public MatchaChildView createView(Context context, MatchaViewNode node) {
                return new CustomView(context, node);
            }
        });
    }

    public CustomView(Context context, MatchaViewNode node) {
        super(context);
        viewNode = node;

        float ratio = (float)context.getResources().getDisplayMetrics().densityDpi / DisplayMetrics.DENSITY_DEFAULT;
        view = new SwitchCompat(context);
        view.setPadding(0, 0, (int)(7*ratio), 0);
        view.setOnCheckedChangeListener(new CompoundButton.OnCheckedChangeListener() {
            public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
                if (isChecked != checked) {
                    checked = isChecked;
                    CustomViewProto.Event event = CustomViewProto.Event.newBuilder().setValue(isChecked).build();
                    CustomView.this.viewNode.call("OnChange", new GoValue(event.toByteArray()));
                }
            }
        });
        addView(view);
    }

    @Override
    public void setNativeState(byte[] nativeState) {
        super.setNativeState(nativeState);
        try {
            CustomViewProto.View proto = CustomViewProto.View.parseFrom(nativeState);
            checked = proto.getValue();
            view.setChecked(proto.getValue());
            view.setEnabled(proto.getEnabled());
        } catch (InvalidProtocolBufferException e) {
        }
    }
}
