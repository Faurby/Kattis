import java.util.Scanner;

public class Chanukah {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int quantity = sc.nextInt();
        for (int i = 0; i < quantity; i++){
            int number = sc.nextInt();
            int days = sc.nextInt();
            System.out.println(number + " " + ((days*(days+1)/2)+days));
        }
        sc.close();
    }
}