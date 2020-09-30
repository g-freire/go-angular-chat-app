/*
 SINGLETON SERVICE TO HANDLE THE SOCKET STREAM FROM THE CLIENT API
 THIS SERVICE SHOULD BE CONNECTED AT BOOT AN THEN MULTIPLEX
 THE DATA TO ALL SUBSCRIBED OBSERVERS
 */


import { EventEmitter, Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';


type NewType = EventEmitter<any>;

@Injectable({
  providedIn: 'root',
})

 export class TopdriveStreamService {

   private streamSubject$: Subject<any> = new Subject<any>()
   private streamEmitter: EventEmitter<any> = new EventEmitter<any>();;
   public streamObservableResult$: Observable<any>;

  private ws: WebSocket
  public webSocketAdress : "ws://localhost:5000/ws" |
                           "ws://localhost:4444/ws" |
                           'wss://topdrive-api-qa.azurewebsites.net/ws'
  mProfileResults

  constructor(
  ) {
    this.webSocketAdress = "ws://localhost:4444/ws"
    this.streamSubject$ = new Subject<any>();
    this.streamObservableResult$ = this.streamSubject$.asObservable();
    this.connectWS()
  }


  public  connectWS(){
    this.ws = new WebSocket(this.webSocketAdress);
    console.log("#################################")
    console.log("TOP DRIVE STREAM SINGLETON STARTED")
    console.log("#################################")

    this.ws.onopen = event => {
      this.streamEmitter.emit({"type": "open", "data": event});
      this.ws.send("CONNECTION OPENED BY FRONTEND CLIENT")
    }
    this.ws.onmessage = (event) => {
      console.log(event.data)
      // this.streamSubject$.next(event)
    }
    this.ws.onclose = event => {
      this.streamEmitter.emit({"type": "close", "data": event});
    }
  }

  public cancelStreaming(){
    this.ws.send("Cancelling Stream")
    this.streamSubject$.complete()
    console.log("#################################")
    console.log("TOP DRIVE STREAM SINGLETON CANCELED")
    console.log("#################################")
  }

  onDestroy(){
    this.streamSubject$.complete()
    this.ws.send("Closing Stream")
    this.ws.close()
    console.log("#################################")
    console.log("TOP DRIVE STREAM SINGLETON DESTROYED")
    console.log("#################################")
  }
}
