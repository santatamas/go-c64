import { Component } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { MatTableModule } from '@angular/material/table';
import { TelemetryService } from '../services/telemetry.service';
import { TablehelperService } from '../services/tablehelper.service';
import { Telemetry } from '../models/telemetry.model';
import { TableRow } from '../models/tablerow.model';

@Component({
  selector: 'app-cia',
  standalone: true,
  imports: [MatCardModule, MatTableModule],
  template: `
    <mat-card>
      <mat-card-header><mat-card-title>CIA</mat-card-title><mat-card-subtitle>Chip states</mat-card-subtitle></mat-card-header>
      <mat-card-content>
        <table mat-table [dataSource]="dataSource" class="mat-elevation-z8">
          <ng-container matColumnDef="name"><th mat-header-cell *matHeaderCellDef>Property</th><td mat-cell *matCellDef="let el">{{el.Name}}</td></ng-container>
          <ng-container matColumnDef="value"><th mat-header-cell *matHeaderCellDef>Value</th><td mat-cell *matCellDef="let el">{{el.Value}}</td></ng-container>
          <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>
          <tr mat-row *matRowDef="let row; columns: displayedColumns;"></tr>
        </table>
      </mat-card-content>
    </mat-card>
  `
})
export class CiaComponent {
  displayedColumns = ['name', 'value'];
  dataSource: TableRow[] = [];

  constructor(private telemetryService: TelemetryService, private helper: TablehelperService) {
    telemetryService.getTelemetry().subscribe((t: string) => {
      const telemetry: Telemetry = JSON.parse(t);
      if (telemetry.Command === 'GetCIAState') {
        this.dataSource = helper.convertToTableRowsWithBlackList(JSON.parse(atob(telemetry.Payload)), ['Keyboard_matrix']);
      }
    });
  }
}
