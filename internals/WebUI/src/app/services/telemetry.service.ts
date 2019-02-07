import { Injectable } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { Observable, Subscriber, observable } from 'rxjs';
import { map, filter, scan } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class TelemetryService {

  observableSocket: Observable<any>;

  constructor(private webSocket: WebsocketService) {
    const openSubscriber = Subscriber.create(() => console.log('connection opened'));
    this.observableSocket = this.webSocket.createObservableSocket('ws://localhost:8000', openSubscriber).pipe(
      map(message => JSON.parse(message)));
  }

  getTelemetry(): Observable<any> {
    return this.observableSocket;
  }

  sendCommand(command: string) {
    this.webSocket.send(command);
  }
}
