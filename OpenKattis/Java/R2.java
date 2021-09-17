import java.util.Scanner;

public class R2{
    public static void main(String[] args) {
    //scanner object
        Scanner sc = new Scanner(System.in);
        String input1 = sc.nextLine();
        sc.close();

        String[] sp = input1.split(" ");

        int r1 = Integer.parseInt(sp[0].trim());
        int s = Integer.parseInt(sp[1].trim());

        int r2 = (s * 2) - r1;

        System.out.println(r2);
        

    }
}