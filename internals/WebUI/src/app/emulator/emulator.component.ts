import { Component, OnInit } from '@angular/core';
import { TelemetryService } from '../services/telemetry.service';
import { Telemetry } from '../models/telemetry.model';
import { EmulatorState } from '../models/emulatorstate.model';
import { TablehelperService } from '../services/tablehelper.service';

@Component({
  selector: 'app-emulator',
  templateUrl: './emulator.component.html',
  styleUrls: ['./emulator.component.scss']
})
export class EmulatorComponent implements OnInit {

  public displayedColumns: string[] = ['name', 'value'];
  public dataSource: any;

  constructor(private telemetryService: TelemetryService, private helper: TablehelperService) {

    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);

      if (telemetry.Command === 'GetEmulatorState') {
        this.dataSource = helper.convertToTableRows(JSON.parse(atob(telemetry.Payload)));
      }
    });
   }

  ngOnInit() {
  }
}
