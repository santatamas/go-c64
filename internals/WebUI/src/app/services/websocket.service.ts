import { Injectable } from '@angular/core';
import { Subscriber, Observable, Subject } from 'rxjs';
import { Telemetry } from '../models/telemetry.model';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {

  private ws: WebSocket;
  private observable: Observable<any>;
  private subject: Subject<any> = new Subject<any>();

  createObservableSocket(url: string, openSubscriber: Subscriber<any>): Subject<any> {
    this.ws = new WebSocket(url);
    new Observable(observer => {
      this.ws.onmessage = event => observer.next(event.data);
      this.ws.onerror = event => observer.error(event);
      this.ws.onclose = event => observer.complete();
      this.ws.onopen = event => {
        openSubscriber.next();
        openSubscriber.complete();
      };

      return () => this.ws.close();
    }).subscribe((data) => { this.subject.next(data); console.log(data); });

    return this.subject;
  }

  send(message: any) {
    this.ws.send(message);
  }
}
