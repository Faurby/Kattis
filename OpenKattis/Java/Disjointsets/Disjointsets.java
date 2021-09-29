import java.util.ArrayList;
import java.util.Scanner;

public class Disjointsets {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        int m = sc.nextInt();
        ArrayList<ArrayList<Integer>> outerList = new ArrayList<ArrayList<Integer>>();
        for(int i = 0; i < n; i++){
            ArrayList<Integer> list = new ArrayList<Integer>();
            list.add(i);
            outerList.add(list);
        }
        
        int S = 0;
        int T = 0;

        for(int i = 0; i < m; i++){
        int operation = sc.nextInt();
        int s = sc.nextInt();
        int t = sc.nextInt();
        
        for(int j = 0; j < outerList.size(); j++){
            if (outerList.get(j).contains(s)){
                S = j;
            }
            if (outerList.get(j).contains(t)){
                T = j;
            }
        }

        if(operation == 0){
            if(S == T){
                System.out.println(1);
            } else {
                System.out.println(0);
            }
        } else if (operation == 1){
            if(S != T){
                outerList.get(S).addAll(outerList.get(T));
                outerList.remove(T);                     
            }
        } else if (operation == 2){
                if(S != T){
                    int index = outerList.get(S).indexOf(s);
                    outerList.get(S).remove(index);
                    outerList.get(T).add(s);
             }
        }
    }
    
    }
}
