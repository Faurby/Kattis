import java.util.Scanner;

public class Pot{
    public static void main(String[] args) {
    //scanner object

        int sum = 0;
        int input;
        int inputPeriods;
        String inputToString;
        int power;
        int number;

        Scanner sc = new Scanner(System.in);
        inputPeriods = sc.nextInt();
 

        for(int i = 0; i < inputPeriods; i++) {
            input = sc.nextInt();
            power = input % 10;
            inputToString = Integer.toString(input);
            number = Integer.parseInt(inputToString.replaceFirst(".$", ""));
            

            sum += Math.pow(number,power);
        }
        
        sc.close();

        System.out.println(sum);

    }
}