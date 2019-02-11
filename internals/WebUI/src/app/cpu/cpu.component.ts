import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { CPUState } from '../models/cpustate.model';

@Component({
  selector: 'app-cpu',
  templateUrl: './cpu.component.html',
  styleUrls: ['./cpu.component.scss']
})
export class CPUComponent implements OnInit {

  public state: CPUState;
  public displayedColumns: string[] = ['register', 'value'];

  constructor(private telemetryService: TelemetryService) {
    this.state = new CPUState();
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);
      if (telemetry.Command === '"GetCPUState"') {
        this.state = JSON.parse(atob(telemetry.Payload));
      }
    });
   }

  ngOnInit() {
  }

}
