import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { CPUState } from '../models/cpustate.model';
import { TableRow } from '../models/tablerow.model';
import { TablehelperService } from '../services/tablehelper.service';

@Component({
  selector: 'app-cpu',
  templateUrl: './cpu.component.html',
  styleUrls: ['./cpu.component.scss']
})
export class CPUComponent implements OnInit {

  public displayedColumns: string[] = ['name', 'value'];
  public dataSource: any;

  constructor(private telemetryService: TelemetryService, private helper: TablehelperService) {
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);
      if (telemetry.Command === 'GetCPUState') {
        this.dataSource = helper.convertToTableRows(JSON.parse(atob(telemetry.Payload)));
      }
    });
   }

  ngOnInit() {
  }

}
