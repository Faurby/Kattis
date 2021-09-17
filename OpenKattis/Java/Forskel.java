import java.util.Scanner;

public class Forskel{
    public static void main(String[] args) {
    //scanner object
        
        
        Scanner sc = new Scanner(System.in);
       
        while(sc.hasNextLong()){
            /*String input1 = sc.nextLine();

            String[] sp = input1.split(" ");

            if (!input1.equals("")) {
                int heltal1 = Integer.parseInt(sp[0]);
                int heltal2 = Integer.parseInt(sp[1]);

                System.out.println(Math.abs(heltal1 - heltal2));
            }*/


            Long firstNumber = sc.nextLong();
            Long secondNumber = sc.nextLong();
            System.out.println(Math.abs(firstNumber - secondNumber));
        } sc.close();
        
        

    }
}