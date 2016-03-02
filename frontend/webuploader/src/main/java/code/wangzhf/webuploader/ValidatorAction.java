package code.wangzhf.webuploader;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class ValidatorAction extends HttpServlet {

	private static final long serialVersionUID = 1L;

	public ValidatorAction() {
		super();
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doPost(request, response);
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		
		String fileName = request.getParameter("fileName");
		if(fileName != null && !"".equals(fileName)){
			if(fileName.indexOf("q") > -1){
				response.getWriter().write("{\"success\": false}");
				System.out.println("重复了");
				return;
			}
		}
		response.getWriter().write("{\"success\": true}");
	}

}
