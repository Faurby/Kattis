public class Party implements Comparable<Party>{
	
	private int seats;
	private int votes;
	
	public Party(int votes){
		this.votes = votes;
	}
	
	public void setSeats(int seats){
		this.seats = seats;
	}
	
	public int getSeats(){
		return seats;
	}
	
	public int getVotes(){
		return votes;
	}
	
	public double getQuotient(){
		return (double) votes / (double) (seats+1);
	}
	
	@Override
	public int compareTo(Party p){
		if(p.getQuotient() < getQuotient()){
			return 1;
		} else if (p.getQuotient() > getQuotient()){
			return -1;
		} 
		return 0;
	}
}