import {
	Component,
	ElementRef,
	EventEmitter,
	inject,
	input,
	Input,
	InputSignal,
	OnChanges,
	OnDestroy,
	OnInit,
	Output,
	SimpleChanges,
	ViewChild,
} from '@angular/core';
import {Subject} from 'rxjs';
import {TranslocoModule} from '@ngneat/transloco';
import {MatIcon} from '@angular/material/icon';
import {NgClass, NgOptimizedImage, NgStyle} from '@angular/common';
import {FormControl, ReactiveFormsModule} from '@angular/forms';
import {environment} from '@environments/environment';
import {Image} from '@app/shared/interfaces/image.interface';
import {HttpClient} from '@angular/common/http';

const ALLOWED_FILE_TYPES = ['image/jpeg', 'image/png'];

const kbToBytes = (kb: number): number => kb * 1024;
const bytesToKb = (bytes: number): number => bytes / 1024;

const constraints = {
	maxFileSize: kbToBytes(300),
	totalFiles: 1,
};

@Component({
	selector: 'image-picker',
	templateUrl: './image-picker.component.html',
	standalone: true,
	imports: [MatIcon, NgClass, TranslocoModule, NgOptimizedImage, NgStyle, ReactiveFormsModule],
})
export class ImagePickerComponent implements OnChanges, OnDestroy, OnInit {
	private http = inject(HttpClient);

	@Input() mandatory: boolean = true;
	@Input() image: string | Image;
	@Input() disabled: boolean = false;
	@Input() required: boolean = false;
	@Input() miniView: boolean = false;
	$isLoading: InputSignal<boolean> = input<boolean>();
	@Output() imageOutput = new EventEmitter();
	@Output() error = new EventEmitter();

	@ViewChild('fileInput', {static: false}) fileInput!: ElementRef;
	allowedFileTypes: string[] = ALLOWED_FILE_TYPES;
	fileName: string | null;
	imageUuid: string;
	fileUrl!: string | null;
	uploadFile!: File | null;
	fileSize!: string | null;

	control: FormControl = new FormControl();
	unsubscribe$: Subject<void> = new Subject<void>();

	ngOnInit(): void {
		if (this.image && typeof this.image === 'string') {
			this.fileUrl = this.image;
		}
	}

	handleChange(event: any) {
		const file = event.target.files[0] as File;
		if (file.size > 1024 * 1024 * 4.5) {
			this.error.emit('File size is too big');
			return;
		}
		this.fileUrl = URL.createObjectURL(file);
		this.uploadFile = file;
		this.fileSize = bytesToKb(file.size).toFixed(2) + ' KB';
		this.fileName = file.name;
		this.imageOutput.emit(file);
	}

	handleRemovesFile() {
		if (this.fileInput && this.fileInput.nativeElement) {
			this.fileInput.nativeElement.value = null;
		}
		this.fileName = null;
		this.uploadFile = null;
		this.fileUrl = null;
		this.imageOutput.emit(null);
	}

	ngOnChanges(changes: SimpleChanges): void {
		if (changes.image && changes.image.currentValue) {
			this.image = changes.image.currentValue;
			this.fileUrl = this.image as string;
			this.imageUuid = this.fileUrl.split('/')[4];
			this.getImageInfo();
		}
		if (changes.submitted) {
			this.handleRemovesFile();
		}
	}

	ngOnDestroy() {
		this.unsubscribe$.next();
		this.unsubscribe$.complete();
	}

	getImageInfo(): void {
		this.http
			.get(`${environment.host + environment.image.host + environment.image.path}${this.imageUuid}`)
			.subscribe((res: any) => {
				this.fileName = res['Name'];
				this.fileSize = bytesToKb(res['Size']).toFixed(2) + ' KB';
			});
	}
}
