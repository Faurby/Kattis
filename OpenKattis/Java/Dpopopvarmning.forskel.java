import java.util.Scanner;

public class DpopopvarmningForskel{
    public static void main(String[] args) {
    //scanner object
        Scanner sc = new Scanner(System.in);
       
        int i = 0;

        while(sc.hasNextLine()){
            String input1 = sc.nextLine();

            String[] sp = input1.split(" ");

            int heltal1 = Integer.parseInt(sp[0].trim());
            int heltal2 = Integer.parseInt(sp[1].trim());

            System.out.println(heltal1 - heltal2);

        } sc.close();
        
        

    }
}