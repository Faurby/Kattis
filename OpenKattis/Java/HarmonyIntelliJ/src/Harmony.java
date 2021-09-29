import java.util.Scanner;

public class Harmony{
	public static void main(String[] args){
			Scanner sc = new Scanner(System.in);
			int vertices = sc.nextInt();
			int edges = sc.nextInt();

			GraphCustom graph = new GraphCustom(vertices);
			
			for(int i = 0; i < edges; i++){
				graph.addEdge(sc.nextInt(), sc.nextInt(), sc.nextInt());
			}
			
			BipartiteXCustom bi = new BipartiteXCustom(graph);
			if(bi.isBipartite()){
				System.out.println(1);
			} else {
				System.out.println(0);
			}
			
	}
}