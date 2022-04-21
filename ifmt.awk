# format *.go file imports
# usage: find . -name "*.go" | xargs -n1 -I {} awk -f ifmt.awk {}

BEGIN {
    TRUE=1
    FALSE=0

    startNum=-1
    endNum=-1
    startFlag=0
    endFlag=0

    local="github.com/hauson/nsq-study"
}

$0 ~ /^import \(/{
    startFlag=1
    startNum=NR
    next
}

startFlag==1 && endFlag==0{
    if (startFlag==1 && $0 ~ /^)$/) {
        startFlag=0
        endFlag=1
        endNum=NR-1
        next
    }

    if ($0 == "") {
        next
    }

    if (isLocal($0)) {
        localPkg[NR]=$0
    } else if (isOfficial($0)) {
        officialPkg[NR]=$0
    } else {
        thirdPkg[NR]=$0
    }
}

END {
    if (startNum < endNum) {
        cmd=sprintf("gsed -i '%d,%dd' %s", startNum+1, endNum,FILENAME)
        system(cmd)
    }

    for (k in officialPkg) {
        cmd=sprintf("gsed -i '%da\\%s' %s", startNum, officialPkg[k], FILENAME)
        system(cmd)
        startNum+=1
    }

    if (isEmpty(officialPkg) == 0 && isEmpty(thirdPkg) == 0 ) {
        cmd=sprintf("gsed -i '%da \n' %s", startNum, FILENAME)
        system(cmd)
        startNum+=1
    }

    for (k in thirdPkg) {
        cmd=sprintf("gsed -i '%da\\%s' %s", startNum, thirdPkg[k], FILENAME)
        system(cmd)
        startNum+=1
    }

    if ((isEmpty(officialPkg)==0 || isEmpty(thirdPkg)==0) && isEmpty(localPkg)==0) {
        cmd=sprintf("gsed -i '%da \n' %s", startNum, FILENAME)
        system(cmd)
        startNum+=1
    }

    for (k in localPkg) {
        cmd=sprintf("gsed -i '%da\\%s' %s", startNum, localPkg[k], FILENAME)
        system(cmd)
        startNum+=1
    }

    cmd=sprintf("go fmt %s", FILENAME)
    system(cmd)
}

func isOfficial(str) {
    if (index(str,".") == 0) {
        return TRUE
    }

    return FALSE
}

func isLocal(str) {
    if (index(str,local) > 0) {
        return TRUE
    } else {
        return FALSE
    }
}

func isEmpty(arr) {
    ret = TRUE
    for (k in arr) {
        ret=FALSE
        break
    }
    return ret
}


