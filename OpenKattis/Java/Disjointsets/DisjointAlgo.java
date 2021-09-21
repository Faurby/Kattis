import java.util.Scanner;

public class DisjointAlgo {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        int m = sc.nextInt();

        UF uf = new UF(n);

        for(int i  = 0; i < m; i++){
            int operation = sc.nextInt();
            int s = sc.nextInt();
            int t = sc.nextInt();
            if(operation  == 0){
                if(uf.connected(s, t)){
                    System.out.println(1);
                } else {
                    System.out.println(0);
                }
            } else if (operation == 1){
                uf.union(s, t);
                } else if (operation == 2) {
                if(!uf.connected(s, t)){
                    uf.move(s, t);
                }
            } 
        }
    }
}