import { Injectable } from '@angular/core';
import { Subscriber, Observable } from 'rxjs';

const SERVER_URL = 'http://localhost:8080';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {

  private ws: WebSocket;

  createObservableSocket(url: string, openSubscriber: Subscriber<any>): Observable<any> {
    this.ws = new WebSocket(url);
    return new Observable(observer => {
      this.ws.onmessage = event => observer.next(event.data);
      this.ws.onerror = event => observer.error(event);
      this.ws.onclose = event => observer.complete();
      this.ws.onopen = event => {
        openSubscriber.next();
        openSubscriber.complete();
      };

      return () => this.ws.close();
    });
  }

  send(message: any) {
    this.ws.send(JSON.stringify(message));
  }
}
