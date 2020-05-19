#!/bin/bash

echo "Pod creation on process..."
sleep 1
for i in {1..10};
do
  kubectl run adhoc$i --image=alpine --generator=run-pod/v1 --restart Never --command -- ping 8.8.8.8
done
