package main
import "fmt"
func main(){
	var ConferenceName="Ease My Trip" //Declaring a variable or a constant
	const ConferenceTickets=50 //If we Declare Const we cannot Change it in later it is final
	var RemainingTickets=50

	fmt.Println("Welcome to ",ConferenceName, " Booking Application");
	fmt.Println("Get your Tickets");
	fmt.Println("Total Tickets",ConferenceTickets,"Available Tickets",RemainingTickets);
	fmt.Printf("Total Tickets %v Available Tickets %v",ConferenceTickets,RemainingTickets);//we use printf when format Specifier is present


}