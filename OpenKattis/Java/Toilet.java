import java.util.Scanner;

public class Toilet{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        char[] sequence = sc.nextLine().toCharArray();
        System.out.println(policyOne(sequence));
        System.out.println(policyTwo(sequence));

        }
        public static int policyOne(char[] list){
            int counter = 0;
            boolean up = false;
            if (list[0] == 'U'){
                up  = true;
            } 

            for (int i = 0; i<list.length ; i++){
                if (list[i] == 'U'){
                    if (up == false){
                        up = true;
                        counter++;
                    } 
                        } else {
                            if (up == true){
                                counter += 2;
                    } else {
                        up = true;
                        counter++;
                    }
                }
            }

            return counter;
    }

    public static int policyTwo(char[] list){
        int counter = 0;
        boolean up = true;
        if (list[0] == 'D'){
            up  = false;
        } 

        for (int i = 0; i<list.length ; i++){
            if (list[i] == 'D'){
                if (up == true){
                    up = false;
                    counter++;
                } 
                    } else {
                        if (up == false){
                            counter += 2;
                } else {
                    up = false;
                    counter++;}
            }
        }
        return counter;
    }
}