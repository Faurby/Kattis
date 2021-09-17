import java.util.Scanner;

public class IsItHalloween{
    public static void main(String[] args) {
    //scanner object
    Scanner sc = new Scanner(System.in);
    
    String date = sc.nextLine();
    sc.close();
    if(date.equals("OCT 31") || date.equals("DEC 25")){
        System.out.println("yup");
    } else {
        System.out.println("nope");
    }

    }
}