public class Student implements Comparable<Student>{
	
	private String grade;
	private String name;
	private String letter;
	private String modifier;
	
	public Student(String name, String grade){
		this.name = name;
		this.grade = grade;
		if(grade.contains("FX")){
			this.letter = grade.substring(0,2);
		} else {
			this.letter = grade.substring(0,1);
		}
		this.modifier = grade.substring(letter.length());
	}
	
	public String getName(){
		return name;
	}
	
	public String getGrade(){
		return grade;
	}
	
	@Override
	public int compareTo(Student studentToCompareTo){
		
		if(studentToCompareTo.getGrade().equals(getGrade())){
			if(studentToCompareTo.getName().compareTo(getName()) < 0){
				return 1;
			} else if (studentToCompareTo.getName().compareTo(getName()) > 0){
				return -1;
			}
		} else {
			if(studentToCompareTo.getGradeNumber()<getGradeNumber()){
				return -1;
			} else if (studentToCompareTo.getGradeNumber()>getGradeNumber()){
				return 1;
			} else {
				if(studentToCompareTo.getModifier()<getModifier()){
					return -1;
				} else if (studentToCompareTo.getModifier()>getModifier()){
					return 1;
				}
			}
		}
		return 0;
	}
	
	public int getModifier(){
		int sum = 0;
		if(modifier.contains("+")){
			sum += modifier.length();
		} else {
			sum -= modifier.length();
		}
		return sum;
	}
	
	public int getGradeNumber(){
		switch(letter){
			case "A":
				return 7;
			case "B":
				return 6;
			case "C":
				return 5;
			case "D":
				return 4;
			case "E":
				return 3;
			case "FX":
				return 2;	
		}
		return 1;
	}
}