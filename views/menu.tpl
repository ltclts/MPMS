<nav id="global_menu" class="menu" data-ride="menu">
    <ul id="treeMenu" class="tree tree-menu" data-ride="tree">
        {{range $i, $elem := .MenuList}}
            <li class="{{if $elem.Active}}open{{end}}">
                <a href="javascript:void(0);"><i class="icon icon-{{$elem.Icon}}"></i>{{$elem.Name}}</a>
                {{range $j, $route := $elem.Routes}}
                    <ul class="nav">
                        <li {{if $route.Active}} class="active"{{end}}>
                            <a href="{{$route.Url}}">&nbsp;&nbsp;{{$route.Name}}</a>
                        </li>
                    </ul>
                {{end}}
            </li>
        {{end}}
</nav>