# for each casio for
mkdir -p casio_decompiled
for file in $(ls casio/); do
    filename=$(echo $file|cut -d. -f1)
    newFile=${filename}.decompiled.g1m
    echo "Decompile $file to casio_decompiled/$newFile"
    echo casio/$file | ./g1mdecompiler > casio_decompiled/$newFile 
done