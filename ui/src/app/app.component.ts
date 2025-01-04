import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CoreModule } from '@onega-ui/core';
import { KitModule } from '@onega-ui/kit';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  standalone: true,
  imports: [CoreModule, KitModule, RouterOutlet],
})
export class AppComponent {
}
