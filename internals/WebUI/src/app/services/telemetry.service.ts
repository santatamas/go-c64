import { Injectable } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { Observable, Subscriber, observable, Subject } from 'rxjs';
import { map, filter, scan } from 'rxjs/operators';
import { Telemetry } from '../models/telemetry.model';

@Injectable({
  providedIn: 'root'
})
export class TelemetryService {

  observableSocket: Subject<any>;

  constructor(private webSocket: WebsocketService) {
    const openSubscriber = Subscriber.create(() => console.log('connection opened'));
    this.observableSocket = this.webSocket.createObservableSocket('ws://localhost:8080/ws', openSubscriber);
  }
  getTelemetry(): Subject<any> {
    return this.observableSocket;
  }

  sendStringCommand(command: string) {
    const telemetryRequest = new Telemetry();
    telemetryRequest.Command = command;
    this.sendCommand(telemetryRequest);
  }

  sendCommand(request: Telemetry) {
    this.webSocket.send(JSON.stringify(request));
  }
}
