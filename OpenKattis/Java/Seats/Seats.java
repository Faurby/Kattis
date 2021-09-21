import java.util.Scanner;
import edu.princeton.cs.algs4.MaxPQ;

public class Seats{
	public static void main(String[] args){ //with help from my friend Thor Tudal. We had a great time
		Scanner sc = new Scanner(System.in);
		
		int parties = sc.nextInt();
		int seats = sc.nextInt();
		
		MaxPQ queue = new MaxPQ();
		
		Party[] array = new Party[parties];
		
		for(int i = 0; i < parties; i++){
			Party p = new Party(sc.nextInt());
			array[i] = p;
			queue.insert(p);
		}
		
		for(int i = 0; i < seats; i++){
			Party p = (Party) queue.delMax();
			p.setSeats(p.getSeats()+1);
			queue.insert(p);
		}
		
		for(int i = 0; i < parties; i++){
			System.out.println(array[i].getSeats());
		}
	}
}