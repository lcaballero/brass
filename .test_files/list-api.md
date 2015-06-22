# List API

Bank :
  name
  createdOn
  updateOn
  rolls : {
    name1: []
    name2: []
    name3: []
  }
  

RollDef :
  Name : string                   // Human name
  Type : [default: ""] string     // External client specific type name.
  Size : int                      // Initial/defaults size
  bank : *Bank                    // Pointer to the containing Bank
  id   : int                      // Internal Bank ID used for possible optimizations
  
/*
  Create Account
  Login
  Dashboard View
  Planned View
  Unfinished View
  Backlog View
  Archive View
*/

/*
  Registers RollDef(s) that def will create a Roll internally in the Bank.
  The returned pointer to a Roll can then be used to access the underlying
  data instead of a constant list-name.
  
  Registration occurs after load or creation of a new Bank.  Internally,
  a Roll is allocated when the bank is loaded, and shared once a registration
  for the correct name (or ID) is submitted via register.
*/
Register( RollDef ) *Roll

NewBank( bank-name )            // [DONE]
LoadBank( bank-name )           // [DONE]

Roll.Bank() *Bank               // Access Pointer to the contain bank
Roll.Add( T... )                // [DONE]
Roll.Insert( index, T )         // [DONE]
Roll.Clear()                    // [DONE]
Roll.Contains( Roll, T, fn )    // [DONE]
Roll.Length()                   // [DONE]

Roll.All( T... ) : bool         // [DONE]
Roll.Any( T... ) : bool         // [DONE]
Roll.IsEmpty() : bool           // [DONE]
Roll.Map( Mapper ) : []T        // [DONE]
Roll.Find( Pred ) : T           // [DONE]
Roll.Where( Pred ) : []T        // [DONE]
 
Roll.Equals( List )
Roll.Get( index )
Roll.IndexOf( T )
Roll.Iterator()
Roll.LastIndexOf( T )
Roll.Remove( T )
Roll.Remove( int )
Roll.RemoveAll( T... )
Roll.RetainAll( T... )
Roll.Set( index, T )
Roll.Sublist( start, end )



