<nav id="global_menu" class="menu" data-toggle="menu">
    <ul class="nav nav-primary">
    {{range $i, $elem := .MenuList}}
        <li class="nav-parent {{if $elem.Active}}show{{end}}">
            <a href="javascript:void(0);"><i class="icon icon-{{$elem.Icon}}"></i>{{$elem.Name}}</a>
        {{range $j, $route := $elem.Routes}}
            <ul class="nav">
                <li {{if $route.Active}} class="active"{{end}}>
                    <a href="{{$route.Url}}">{{$route.Name}}</a>
                </li>
            </ul>
        {{end}}
        </li>
    {{end}}
    </ul>
</nav>
