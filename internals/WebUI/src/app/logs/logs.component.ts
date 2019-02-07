import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';

@Component({
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.scss']
})
export class LogsComponent implements OnInit {

  latestMessage: string;

  constructor(telemetryService: TelemetryService) {
    telemetryService.getTelemetry().subscribe((t: any) => {
      //console.log(t);
      console.log("message received");
      this.latestMessage = t;
    });
   }

  ngOnInit() {
  }

}
