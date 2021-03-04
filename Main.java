import components.set.Set;
import components.set.Set1L;

public class Main {

    public static void main(String[] args) {
        Set<String> test = new Set1L<String>();
        test.add("Test");
        System.out.println(test);
    }
}