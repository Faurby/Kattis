import edu.princeton.cs.algs4.UF;
import java.util.Scanner;

public class Tildes {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        int m = sc.nextInt();
         
        UF uf = new UF(n+1);
		int a = 0;
		int b = 0;
        for(int i  = 0; i < m; i++){
            String operation = sc.next();
			if(operation.equals("s")){
				a = sc.nextInt();
				int counter = 1;
				for(int j = 0; j < uf.count(); j++){
					if(b != uf.find(a)){
						counter++;
						b = uf.find(a);
					}
					
				}
				System.out.println(counter);
			} else if (operation.equals("t")){
				a = sc.nextInt();
				b = sc.nextInt();
				uf.union(a, b);
				System.out.println("a: " + a + "uf: " + uf.find(a) + "b: " + b + "uf: " + uf.find(b));
			}
			
            
        }
    }
}