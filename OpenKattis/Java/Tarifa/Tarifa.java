import java.util.Scanner;

public class Tarifa {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int megabytes = sc.nextInt();
        int months = sc.nextInt();
        int sum = 0;
        for (int i = 0; i < months; i++) {
            sum += sc.nextInt();
        }
        System.out.println((megabytes* (months + 1) - sum));
    }
}
