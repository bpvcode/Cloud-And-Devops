import {
  Component,
  ChangeDetectorRef
} from '@angular/core';

import { fadeAnimation } from '../app/shared/animations/fade-in.animation';

import * as typer from 'typer-js';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  animations: [fadeAnimation],
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'template';
  myimage:string ="assets/background.jpeg";
  showTechnologies: boolean = true;

  private readonly cursor: any;
  public showTechnologies_: boolean;

  constructor( private cdr: ChangeDetectorRef) { }

  public ngAfterViewInit(): void {
    typer('.presentation h1', 90)
      .cursor(this.cursor)
      .line({ container: '.name' })
      .pause(2000)
      .end(() => this.writeTitleAndTechnologies());
  }

  private writeTitleAndTechnologies(): void {
    // Force the re-render of the technologies heading because typer
    // messes with the DOM, erasing (not just hiding) the spans
    this.showTechnologies = false;
    this.cdr.detectChanges();
    this.showTechnologies = true;
    this.cdr.detectChanges();
    typer('.presentation h2', 90)
      .cursor(this.cursor)
      .line({ container: '.title' })
      .pause(2000)
      .continue({ container: '.dot' })
      .continue({ container: '.jscript' })
      .pause(1500)
      .back('empty', -21)
      .continue({ container: '.style' })
      .pause(1500)
      .back('empty', -21)
      .continue({ container: '.framework' })
      .pause(1500)
      .back('empty', -21)
      .continue({ container: '.markup' })
      .pause(2000)
      .back('empty', -21)
      .continue({ container: '.backend' })
      .pause(1500)
      .back('empty', -21)
      .continue({ container: '.db' })
      .pause(1500)
      .back('empty', -21)
      .continue({ container: '.devops' })
      .pause(2000)
      .back('empty', -21)
      .continue({ container: '.obs' })
      .pause(5000)
      .back('empty', -21)
      .empty()
      .end(() => this.writeTitleAndTechnologies());
  }
}



