#!/bin/bash
plugdir=p
if [ -d $plugdir ]; then
    rm $plugdir/*
else
    mkdir $plugdir
fi

pwd=`pwd`
for dir in ./plugins/*/
do
    dir=${dir%*/}
    plugin_name=${dir##*/}
    echo building plugin: ${plugin_name}
    cd ${dir}
    go build --buildmode=plugin -o ${pwd}/${plugdir}/${plugin_name}
    cd ${pwd}
done
echo done! plugins are in ./${plugdir}
