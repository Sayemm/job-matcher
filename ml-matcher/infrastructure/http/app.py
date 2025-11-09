from flask import Flask, request, jsonify
from flask_cors import CORS
import PyPDF2
import io

class MatcherAPI:
    """Infrastructure - HTTP Server using Flask"""
    
    def __init__(self, match_resume_use_case):
        self.use_case = match_resume_use_case
        self.app = Flask(__name__)
        CORS(self.app)
        self._setup_routes()
    
    def _setup_routes(self):
        """Register HTTP endpoints"""
        
        @self.app.route('/health', methods=['GET'])
        def health():
            """Health check endpoint"""
            return jsonify({
                'status': 'healthy',
                'service': 'ml-matcher'
            }), 200
        
        @self.app.route('/parse-pdf', methods=['POST'])
        def parse_pdf():
            """
            Parse PDF and extract text
            
            Request: PDF file in body
            Response: {"text": "extracted text..."}
            """
            try:
                # Get PDF data from request
                pdf_data = request.get_data()
                
                if not pdf_data:
                    return jsonify({'error': 'No PDF data provided'}), 400
                
                # Parse PDF
                pdf_file = io.BytesIO(pdf_data)
                pdf_reader = PyPDF2.PdfReader(pdf_file)
                
                # Extract text from all pages
                text = ""
                for page in pdf_reader.pages:
                    text += page.extract_text() + "\n"
                
                if not text.strip():
                    return jsonify({'error': 'No text found in PDF'}), 400
                
                return jsonify({'text': text}), 200
                
            except Exception as e:
                return jsonify({'error': f'Failed to parse PDF: {str(e)}'}), 500
        
        @self.app.route('/match', methods=['POST'])
        def match():
            """
            Match resume to cluster
            
            Request: {"resume_text": "..."}
            Response: {"cluster_id": 5, "score": 0.87}
            """
            try:
                data = request.get_json()
                
                if not data or 'resume_text' not in data:
                    return jsonify({'error': 'resume_text is required'}), 400
                
                resume_text = data['resume_text']
                
                # Call use case
                result = self.use_case.execute(resume_text)
                
                return jsonify(result), 200
                
            except ValueError as e:
                return jsonify({'error': str(e)}), 400
            except Exception as e:
                return jsonify({'error': f'Internal error: {str(e)}'}), 500
    
    def run(self, host='0.0.0.0', port=5000):
        """Start the HTTP server"""
        print(f"🚀 ML Matcher API starting on {host}:{port}")
        print(f"📍 Endpoints:")
        print(f"   GET  /health")
        print(f"   POST /parse-pdf")
        print(f"   POST /match")
        self.app.run(host=host, port=port)