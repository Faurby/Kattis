import java.util.Scanner;

public class QALY{
    public static void main(String[] args) {
    //scanner object

        float sum = 0;
        String input;
        int inputPeriods;
        float quality;
        float years;

        Scanner sc = new Scanner(System.in);
        inputPeriods = sc.nextInt();
        sc.nextLine();
 

        for(int i = 0; i < inputPeriods; i++) {
            input = sc.nextLine();
            String[] inputSplit = input.split(" ");

            quality = Float.parseFloat(inputSplit[0]);
            years = Float.parseFloat(inputSplit[1]);

            sum += quality * years;
        }
        
        sc.close();

        System.out.println(sum);

    }
}