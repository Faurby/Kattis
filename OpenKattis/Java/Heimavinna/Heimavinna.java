import java.util.Scanner;

public class Heimavinna {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String line = sc.nextLine();
        String[] intervals = line.split(";");
        int sum = 0;
        for (int i = 0; i < intervals.length ; i++) {
            if (intervals[i].contains("-")) {
                 String numbers[] = intervals[i].split("-");
                 int min = Integer.parseInt(numbers[0]);
                 int max = Integer.parseInt(numbers[1]);
                 sum += (max-min) + 1;
            } else {
                sum += 1;
            }
        }
        System.out.println(sum);
    }
}