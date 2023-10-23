// Context is a way for you to add a timeout or a cancellation to a GoRoutine.  
// e.g. Lets say you have a long expensive operation that takes 30 sec to complete and you want to cancel it 
// halfway through then we can use context. 
// On Postman if we make a request to our API and cancel it as we dont need it then context will cancel it
// when client's connection closes and does not execute next logic as user has canceled the request.
