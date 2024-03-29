import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class ArrayUtil {
    /**
     * List去重，不打乱原来顺序，泛型list对象
     * 对象重写hashCode和equals
     * @param <T>
     * @param list
     * @return
     */
    public static <T> List<T> distinctBySetOrder(List<T> list){
        Set<T> set = new HashSet<T>();
        List<T> newList = new ArrayList<T>();
        for(T t: list){
            if(set.add(t)){
                newList.add(t);
            }
        }
        return newList;
    }

    /**
     * List去重，可能打乱原来顺序，泛型list对象
     * 对象重写hashCode和equals
     * @param list
     * @return
     */
    public static <T> List<T> distinctBySet(List<T> list){
        return new ArrayList<T>(new HashSet<T>(list));
    }

    /**
     * list转为字符串，专用于sql中in函数
     * @return String
     * @param list
     */
    public static String strListToSqlJoin(List<String> list) {
        if (null == list || list.size() < 1) {
            return "";
        }
        StringBuilder temp = new StringBuilder();
        for (int i = 0; i < list.size(); i++) {
            if (i > 0) {
                temp.append(",");
            }
            temp.append("'").append(list.get(i)).append("'");
        }
        return temp.toString();
    }
}
