import {Component, Inject, Input, OnInit} from '@angular/core';
import { CommonModule } from '@angular/common';
import {MAT_DIALOG_DATA, MatDialogModule, MatDialogRef} from '@angular/material/dialog';

class DialogData {
    text: string;
    image: string;
    negativeButton: string;
    positiveButton: string;
}

@Component({
  selector: 'app-save-modal',
  standalone: true,
  imports: [CommonModule,MatDialogModule],
  templateUrl: './save-modal.component.html',
  styleUrls: ['./save-modal.component.scss']
})
export class SaveModalComponent implements OnInit {


  constructor(public dialogRef: MatDialogRef<SaveModalComponent>,
      @Inject(MAT_DIALOG_DATA) public data: DialogData) { }

  ngOnInit(): void {
  }

    onNegativeClick(): void {
        this.dialogRef.close();
    }

}
