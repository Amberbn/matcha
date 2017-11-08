package io.gomatcha.bridge;

import java.util.Map;
import java.util.HashMap;
import android.util.Log;
import java.lang.reflect.Method;
import java.lang.reflect.InvocationTargetException;
import java.util.Arrays;

public class Tracker {
    private static final Object testObject = new Object();
    private static final Tracker instance = new Tracker();
    private Map<Long, Object> mapTable = new HashMap<Long, Object>();
    private long maxKey = 0;
    private Tracker() {
    }
    public static Tracker singleton() {
        return instance;
    }
    public synchronized long track(Object v) {
        this.maxKey += 1;
        this.mapTable.put(this.maxKey, v);
        return this.maxKey;
    }
    public synchronized void untrack(long v) {
        if (!this.mapTable.containsKey(v)) {
            throw new IllegalArgumentException("Tracker doesn't contain key");
        }
        this.mapTable.remove(v);
    }
    public synchronized long trackerCount() {
        return this.mapTable.size();
    }
    public synchronized Object get(long v) {
        Object a = this.mapTable.getOrDefault(v, Tracker.testObject);
        if (a == Tracker.testObject) {
            throw new IllegalArgumentException("Tracker doesn't contain key");
        }
        return a;
    }
    public synchronized long foreignBridge(String key) {
        Bridge bridge = Bridge.singleton();
        return track(bridge.get(key));
    }
    public synchronized long foreignCall(long v, String method, long[] args) {
        int len = args.length;
        Object[] vb = new Object[len];
        Class[] vc = new Class[len];
        for (int i = 0; i < len; i++) {
            Object e = this.get(args[i]);
            vb[i] = e;
            vc[i] = e.getClass();
        }
        
        long test = 0;
        try {
            Object a = this.get(v);
            Method m = a.getClass().getMethod(method, vc);
            Object rlt = m.invoke(a, vb);
            test = track(rlt);
        } catch (NoSuchMethodException e) {
            Log.v("Bridge", String.format("foreignCall, %d, %s, %s, %s", v, method, Arrays.toString(args), e.getCause()));
            throw new RuntimeException(e);
        } catch (IllegalAccessException e) {
            Log.v("Bridge", String.format("foreignCall, %d, %s, %s, %s", v, method, Arrays.toString(args), e.getCause()));
            throw new RuntimeException(e);
        } catch (InvocationTargetException e) {
            Log.v("Bridge", String.format("foreignCall, %d, %s, %s, %s", v, method, Arrays.toString(args), e.getCause()));
            throw new RuntimeException(e);
        }
        return test;
    }
    public synchronized long foreignNil() {
        return track(null);
    }
    public synchronized boolean foreignIsNil(long v) {
        Object a = this.get(v);
        return a == null;
    }
    public synchronized long foreignBool(boolean v) {
        return track(v);
    }
    public synchronized boolean foreignToBool(long v) {
        boolean a = (Boolean)this.get(v);
        return a;
    }
    public synchronized long foreignInt64(long v) {
        return track(v);
    }
    public synchronized long foreignToInt64(long v) {
        Object a = this.get(v);
        if (a instanceof Integer) {
            return ((Integer)a).longValue();
        }
        return (Long)a;
    }
    public synchronized long foreignFloat64(double v) {
        return track(v);
    }
    public synchronized double foreignToFloat64(long v) {
        Object a = this.get(v);
        if (a instanceof Float) {
            return ((Float)a).doubleValue();
        }
        return (Double)a;
    }
    public synchronized long foreignGoRef(long v) {
        return track(new GoValue(v, false));
    }
    public synchronized long foreignToGoRef(long v) {
        return ((GoValue)this.get(v)).goRef;
    }
    public synchronized long foreignString(String v) {
        return track(v);
    }
    public synchronized String foreignToString(long v) {
        return (String)this.get(v);
    }
    public synchronized long foreignBytes(byte[] v) {
        return track(v);
    }
    public synchronized byte[] foreignToBytes(long v) {
        return (byte[])this.get(v);
    }
    public synchronized long foreignArray(long[] v) {
        Object[] a = new Object[v.length];
        for (int i = 0; i < v.length; i++) {
            a[i] = this.get(v[i]);
        }
        return track(a);
    }
    public synchronized long[] foreignToArray(long v) {
        Object[] a = (Object[])this.get(v);
        long[] fgnRefs = new long[a.length];
        for (int i = 0; i < a.length; i++) {
            fgnRefs[i] = track(a[i]);
        }
        return fgnRefs;
    }
    public synchronized void foreignPanic() {
        throw new RuntimeException("Golang Panic");
    }
}
