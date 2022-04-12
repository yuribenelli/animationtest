import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  //is new Character Session visible?
  isVisible : boolean = false;
  //
  name:string='';
  //label of NEW Button: can be NEW or CLOSE
  newBtnState:string='New';
  headerLabel:string="Saved Character"
  height:string='8vw';


  constructor() { }

  ngOnInit(): void {}

  onClick(){
    if (this.isVisible == false){
      this.isVisible = true
      this.newBtnState ='Close'
      this.headerLabel ="New Character"
      this.height = '30vw'

    }else{
      this.isVisible = false
      this.newBtnState ='New'
      this.headerLabel ='Saved Character'
      this.height = '8vw'
    }
  }
  getDisplay(){
    if (this.isVisible){
      return "block"
    }else{
      return "none"
    }
  }





}
