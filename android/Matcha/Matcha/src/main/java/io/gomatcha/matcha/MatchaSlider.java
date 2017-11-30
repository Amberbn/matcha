package io.gomatcha.matcha;

import android.content.Context;
import android.widget.SeekBar;

import com.google.protobuf.InvalidProtocolBufferException;

import io.gomatcha.bridge.GoValue;
import io.gomatcha.matcha.proto.view.PbSlider;

class MatchaSlider extends MatchaChildView {
    SeekBar view;
    double value;
    double maxValue;
    double minValue;
    MatchaViewNode viewNode;

    static {
        MatchaView.registerView("gomatcha.io/matcha/view/slider", new MatchaView.ViewFactory() {
            @Override
            public MatchaChildView createView(Context context, MatchaViewNode node) {
                return new MatchaSlider(context, node);
            }
        });
    }

    public MatchaSlider(Context context, MatchaViewNode node) {
        super(context);
        viewNode = node;

        view = new SeekBar(context);
        view.setMax(10000);
        view.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int i, boolean b) {
                if (b && i != value ) {
                    double maxValue = MatchaSlider.this.maxValue;
                    double minValue = MatchaSlider.this.minValue;
                    PbSlider.SliderEvent proto = PbSlider.SliderEvent.newBuilder().setValue((double) i / 10000.0 * (maxValue - minValue) + minValue).build();
                    MatchaSlider.this.viewNode.call("OnValueChange", new GoValue(proto.toByteArray()));
                }
            }
            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {

            }
            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {
                double maxValue = MatchaSlider.this.maxValue;
                double minValue = MatchaSlider.this.minValue;
                PbSlider.SliderEvent proto = PbSlider.SliderEvent.newBuilder().setValue((double)MatchaSlider.this.view.getProgress()/10000.0 * (maxValue - minValue) + minValue).build();
                MatchaSlider.this.viewNode.call("OnSubmit", new GoValue(proto.toByteArray()));
            }
        });
        addView(view);
    }

    @Override
    public void setNativeState(byte[] nativeState) {
        super.setNativeState(nativeState);
        try {
            PbSlider.Slider proto = PbSlider.Slider.parseFrom(nativeState);
            view.setEnabled(proto.getEnabled());
            view.setProgress((int)((proto.getValue()- proto.getMinValue())*10000.0/(proto.getMaxValue() - proto.getMinValue())));
            this.value = view.getProgress();
            this.maxValue = proto.getMaxValue();
            this.minValue = proto.getMinValue();
        } catch (InvalidProtocolBufferException e) {
        }
    }
}
