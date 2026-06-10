import { Injectable } from '@angular/core';
import { Subject, Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class WebsocketService {
  private ws!: WebSocket;
  private subject = new Subject<string>();

  connect(url: string): Subject<string> {
    this.ws = new WebSocket(url);
    this.ws.onmessage = (event) => this.subject.next(event.data);
    this.ws.onerror = (event) => console.error('WebSocket error', event);
    this.ws.onclose = () => console.log('WebSocket closed');
    this.ws.onopen = () => console.log('WebSocket connected');
    return this.subject;
  }

  send(message: string): void {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(message);
    }
  }
}
