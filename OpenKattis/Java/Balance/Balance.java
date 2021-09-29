package Balance;

import java.util.Scanner;
import java.util.Stack;

public class Balance { //Me and Thor Tudal collaborated on this exercise and had a great time

    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);

        while (sc.hasNext()) {
            Stack<Character> stack = new Stack<>();
            String str = sc.next();
            int length = str.length();
			boolean stopped = false;
            for (int i = 0; i < length; i++) {
                char c = str.charAt(i);
               
				if (c == '(' || c == '['){
					stack.push(c);
					stopped = true;
				} else {
					if(stack.isEmpty()){
						stopped = true;
						break;
					} else {
						stopped = false;
						char pop = stack.pop();
						if (pop == '(' && c != ')'){
							stopped = true;
							break;
						}
						if (pop == '[' && c != ']'){
							stopped = true;
							break;
						}
					}
				}
            }
			if(!stopped && stack.isEmpty()){
				System.out.println(1);
			} else {
				System.out.println(0);
			}
			
		}
	}    
}