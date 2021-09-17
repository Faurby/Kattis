import java.util.Scanner;

public class FiftyShades {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int N = sc.nextInt();
        sc.nextLine();

        int pinkButtons = 0;

        for (int i = 0; i < N; i++){
            String button = sc.nextLine();
            button = button.toLowerCase();

            if(button.contains("pink") || button.contains("rose")){
                pinkButtons++;

            }
        }

        if (pinkButtons == 0){
            System.out.println("I must watch Star Wars with my daughter");
        } else System.out.println(pinkButtons);
    }
}
