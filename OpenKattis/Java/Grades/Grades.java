import java.util.Scanner;
import edu.princeton.cs.algs4.Insertion;

public class Grades{
	public static void main(String[] args){
		Scanner sc = new Scanner(System.in);
		int limit = sc.nextInt();
		
		
		Student[] array = new Student[limit];
		
		for(int i = 0; i < limit; i++){
			String name = sc.next();
			String grade = sc.next();
			array[i] = new Student(name, grade);
		}
		
		
		Insertion.sort(array);
		
		for (int i = 0; i < array.length; i++){
			System.out.println(array[i].getName());
		}
		
	}
}