package code.wangzhf.webuploader;

import java.io.File;
import java.io.IOException;
import java.util.Date;
import java.util.Iterator;
import java.util.List;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.FileItemFactory;
import org.apache.commons.fileupload.FileUploadException;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;

/**
 * 文件上传操作
 * @author wangzhf
 *
 */
public class FileAction extends HttpServlet {
	private static final long serialVersionUID = 1L;

	public static final String BASE_PATH = "D:/tmp";
	
	public FileAction() {
	}

	protected void doGet(HttpServletRequest request,
			HttpServletResponse response) throws ServletException, IOException {
		doPost(request, response);
	}

	protected void doPost(HttpServletRequest request,
			HttpServletResponse response) throws ServletException, IOException {
		boolean isMulti = ServletFileUpload.isMultipartContent(request);
		if(isMulti){
			try {
				FileItemFactory factory = new DiskFileItemFactory();
				ServletFileUpload upload = new ServletFileUpload(factory);
				List<FileItem> items = upload.parseRequest(request);
				Iterator<FileItem> it = items.iterator();
				int i = 1;
				int j = 0;
				int m = i / j;
				while(it.hasNext()){
					FileItem item = it.next();
					if(item.isFormField()){
						String fieldName = item.getFieldName();
						String value = item.getString();
						if(fieldName.toLowerCase().indexOf("date") > -1){
							value = new Date(value).toLocaleString();
						}
						System.out.println("普通字段名称是：" + fieldName + ", value: " + value);
					}else{
						File dir = new File(BASE_PATH);
						if(!dir.exists()){
							dir.mkdir();
						}
						String fileName = item.getName();
						System.out.println("文件名：" + fileName);
						File saveFile = new File(BASE_PATH, fileName);
						item.write(saveFile);
					}
				}
			} catch (Exception e) {
				e.printStackTrace();
				//response.sendError(100001);
				//response.sendError(1000, "hello world");
				//throw new RuntimeException("heheh");
				response.setCharacterEncoding("UTF-8");
				response.getWriter().write("{\"error\": true, \"msg\": \"文件已存在\"}");
			}
			
		}
		
	}

}
