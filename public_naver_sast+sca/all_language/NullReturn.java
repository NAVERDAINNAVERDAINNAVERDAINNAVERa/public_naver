package nullReturn;

/**
 * Created by BJ on 2015-05-18.
 */
public class NullReturn {
    public NullReturn() {
    }
    public native Object[] returnNull();
    public static void main(String[] args) {
        NullReturn n = new NullReturn();
        Object[] arr = n.returnNull();
        System.out.println("Array length: " + arr.length);
    }
    static {
        System.loadLibrary("NullReturn");
    }
}
