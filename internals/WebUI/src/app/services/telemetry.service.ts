import { Injectable } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { Observable, Subscriber } from 'rxjs';
import { map, filter, scan } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class TelemetryService {

  constructor(private webSocket: WebsocketService) {}

  watchProduct(productId: number): Observable<any> {
    const openSubscriber = Subscriber.create(
        () => this.webSocket.send({productId: productId}));

    return this.webSocket.createObservableSocket('ws://localhost:8000', openSubscriber).pipe(
        map(message => JSON.parse(message)));
  }
}
