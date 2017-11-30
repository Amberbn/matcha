package io.gomatcha.matcha;

import android.animation.Animator;
import android.animation.AnimatorListenerAdapter;
import android.content.Context;
import android.view.View;
import android.widget.RelativeLayout;

import java.util.ArrayList;
import java.util.List;

class MatchaStackView2 extends RelativeLayout {
    public MatchaStackView2(Context context) {
        super(context);
    }

    List<View> childViews = new ArrayList<View>();
    View topView;

    List<View> getChildViews() {
        return childViews;
    }

    void setChildViews(List<View> v) {
        boolean enter = childViews.size() <= v.size();

        childViews = v;

        if (childViews.size() > 0) {
            long duration = getResources().getInteger(android.R.integer.config_shortAnimTime);
            View top = childViews.get(childViews.size()-1);
            if (enter) {
                if (top.getParent() != null) {
                    this.removeView(top);
                }
                this.addView(top);
                if (childViews.size() > 1) {
                    top.setAlpha(0f);
                    top.setTranslationY(500);
                    top.animate()
                            .translationY(0)
                            .alpha(1)
                            .setDuration(duration)
                            .setListener(new AnimatorListenerAdapter() {
                                @Override
                                public void onAnimationEnd(Animator animation) {
                                    reload();
                                }
                            });
                }
            } else {
                if (top.getParent() != null) {
                    this.removeView(top);
                }
                this.addView(top, 0);
                topView.animate()
                        .translationY(500)
                        .alpha(0)
                        .setDuration(duration)
                        .setListener(new AnimatorListenerAdapter() {
                            @Override
                            public void onAnimationEnd(Animator animation) {
                                reload();
                            }
                        });
            }
            topView = top;
        }
    }

    void reload() {
        int count = this.getChildCount();
        for (int i = 0; i < count; i++) {
            View v = this.getChildAt(count-1-i);
            if (v != topView) {
                this.removeView(v);
            }
        }
    }
}
