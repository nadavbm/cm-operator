#!/bin/sh
kubectl delete crd opconfigmaps.opconfigmap.example.com
kubectl delete clusterrolebinding cm-operator
kubectl delete clusterrole cm-operator
kubectl delete ns cm-operator
kubectl delete ns cmop-test