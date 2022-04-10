import { Component, OnInit } from '@angular/core';
import { Character } from 'src/app/model/character.model';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-pgcard',
  templateUrl: './pgcard.component.html',
  styleUrls: ['./pgcard.component.scss']
})
export class PgcardComponent implements OnInit {


  charactersList: Character[]=[];

 private http: HttpClient;
  constructor(h:HttpClient) {
    this.http = h;
   }

  ngOnInit(): void {
    this.refresh();
  }


  refresh():void{
    this.http.get<Character[]>(environment.apiUrl + "character").subscribe(data => this.populate(data));
  }
  private populate(data: Character[]){
    this.charactersList = data;
  }

}
